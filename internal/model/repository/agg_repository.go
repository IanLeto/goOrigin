package repository

import "goOrigin/internal/model/entity"

func AggregateByTraceID(logs []entity.KafkaLogEntity) map[string]entity.TransInfoEntity {
	aggregated := make(map[string]entity.TransInfoEntity)

	for _, log := range logs {
		// 仅当 TraceId 非空时进行聚合
		if log.TraceId == "" {
			continue
		}

		// 取出 Channel、PodName、Project
		transInfo := entity.TransInfoEntity{
			Cluster: log.SysName,        // 取 SysName 作为 Cluster
			PodName: log.ContainerPodID, // 取 PodName
			Project: log.SysName,        // 取 SysName 作为 Project
		}

		// 以 TraceId 作为 key 进行聚合
		aggregated[log.TraceId] = transInfo
	}

	return aggregated
}
