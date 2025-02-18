package repository

import "goOrigin/internal/model/entity"

func AggregateByTraceID(logs []entity.KafkaLogEntity) map[string]entity.TransInfoEntity {
	aggregated := make(map[string]entity.TransInfoEntity)

	for _, log := range logs {
		// 仅当 TraceId 非空时进行聚合
		if log.TraceId == "" {
			continue
		}

		// 取出 Channel、PodName、SvcName
		transInfo := entity.TransInfoEntity{
			Cluster: log.SysName,            // 取 SysName 作为 Cluster
			Channel: log.Trans.TransChannel, // 取 TransChannel 作为 Channel
			PodName: log.ContainerPodID,     // 取 PodName
			SvcName: log.SysName,            // 取 SysName 作为 SvcName
		}

		// 以 TraceId 作为 key 进行聚合
		aggregated[log.TraceId] = transInfo
	}

	return aggregated
}
