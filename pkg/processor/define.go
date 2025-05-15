package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goOrigin/internal/model/entity"
	"goOrigin/pkg/moniter"
	"goOrigin/pkg/utils"
	"os"
	"reflect"
)

var logger = func() *zap.Logger {
	// 配置 zapcore.Encoder (日志格式配置)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",                        // 时间字段的键名
		LevelKey:       "level",                       // 日志级别字段的键名
		NameKey:        "logger",                      // 日志器名称字段的键名
		CallerKey:      "caller",                      // 调用者信息字段的键名
		MessageKey:     "msg",                         // 日志消息字段的键名
		StacktraceKey:  "stacktrace",                  // 堆栈跟踪字段的键名
		LineEnding:     zapcore.DefaultLineEnding,     // 每行日志的结束符
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 日志级别大写编码 (INFO, WARN, ERROR)
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // 时间格式 (ISO8601 格式)
		EncodeDuration: zapcore.StringDurationEncoder, // 时间间隔格式
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 调用者信息格式 (简短路径)
	}

	// 设置日志输出级别 (DEBUG、INFO、WARN、ERROR)
	level := zapcore.DebugLevel // 写死为 DEBUG 级别，可以用配置加载

	// 日志写入目标
	// 1. 输出到标准输出 (控制台)
	consoleWriter := zapcore.Lock(os.Stdout)

	// 2. 输出到文件 (额外配置)
	fileWriter, err := os.Create("app.log") // 写死日志文件路径为 "app.log"
	if err != nil {
		panic("无法创建日志文件: " + err.Error())
	}

	// 创建 zapcore.Core
	core := zapcore.NewTee(
		// 使用 Tee 将日志同时写入多个输出
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleWriter, level),            // 控制台输出
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(fileWriter), level), // 文件输出 (JSON 格式)
	)

	// 构建 zap.Logger 实例
	logger := zap.New(core,
		zap.AddCaller(),                       // 显示调用者信息 (文件名和行号)
		zap.AddCallerSkip(1),                  // 调用栈跳过级别 (适配封装的 logger)
		zap.AddStacktrace(zapcore.ErrorLevel), // 仅在 ERROR 级别记录堆栈
	)

	return logger
}()

type Processor interface {
	Process(ctx context.Context, input chan []byte, out chan []byte)
}

type FilterProcessor struct {
}

func (p *FilterProcessor) Process(ctx context.Context, input chan []byte, out chan []byte) {
	for data := range input {
		var spanInfo *entity.KafkaLogEntity

		utils.JsonToStruct(data, spanInfo)
		v, err := json.Marshal(spanInfo)
		logger.Sugar().Errorln(err)
		out <- v
	}
}

type MetricProcessor struct {
	RetCodes []string
}

func (p *MetricProcessor) Process(ctx context.Context, input chan []byte, out chan []byte) {
	for data := range input {
		// 使用目标结构体解码 JSON
		var spanInfo entity.SpanTransTypeInfoEntity
		if err := json.Unmarshal(data, &spanInfo); err != nil {
			// 如果反序列化失败，跳过当前数据
			fmt.Printf("Failed to unmarshal data: %v\n", err)
			continue
		}

		// 使用反射提取结构体字段作为 Prometheus 标签
		labels := extractLabelsUsingReflection(spanInfo)

		// 获取 ReturnCode
		retCode, ok := labels["ret_code"]
		if ok && utils.IncludeString(p.RetCodes, retCode) {
			// 增加指标计数
			moniter.SpanCount.With(labels).Inc()
		}

		// 将数据传递到下一个处理阶段（如果需要）
		out <- data
	}
}

// 使用反射提取结构体字段作为 Prometheus 标签
func extractLabelsUsingReflection(obj interface{}) prometheus.Labels {
	labels := prometheus.Labels{}

	// 获取对象的值和类型
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	// 确保是结构体
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		fmt.Println("Provided object is not a struct")
		return labels
	}

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 获取字段的 JSON 标签作为标签名
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}

		// 如果字段值是字符串类型
		if value.Kind() == reflect.String {
			labels[tag] = value.String()
		}
	}

	return labels
}

// Node 接口定义了节点的处理方法
type Node interface {
	Process(ctx context.Context, input []byte) ([]byte, error)
	ProcessWithChannel(ctx context.Context, input <-chan []byte, output chan<- []byte) error
}
