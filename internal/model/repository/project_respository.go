package repository

import (
	"goOrigin/internal/model/dao"
	"goOrigin/internal/model/entity"
)

func ToProjectDAO(project *entity.Project) *dao.TProject {
	return &dao.TProject{
		Project:   project.Project,
		ProjectCN: project.ProjectCN,
		AZ:        project.AZ,
		TracePOD:  project.TracePOD,
		CreateAT:  project.CreateAT,
		UpdateAT:  project.UpdateAT,
	}
}

func ToProjectEntity(tProject *dao.TProject) *entity.Project {
	return &entity.Project{
		ID:        tProject.ID,
		Project:   tProject.Project,
		ProjectCN: tProject.ProjectCN,
		AZ:        tProject.AZ,
		TracePOD:  tProject.TracePOD,
		CreateAT:  tProject.CreateAT,
		UpdateAT:  tProject.UpdateAT,
	}
}

func ToProjectBizCodeDAO(projectBizCode *entity.ProjectBizCode) *dao.TProjectBizCode {
	return &dao.TProjectBizCode{
		BizKey:      projectBizCode.BizKey,
		BizValue:    projectBizCode.BizValue,
		BizType:     projectBizCode.BizType,
		Project:     projectBizCode.Project,
		Cluster:     projectBizCode.Cluster,
		Service:     projectBizCode.Service,
		SpanAlias:   projectBizCode.SpanAlias,
		CreatedUser: projectBizCode.CreatedUser,
		UpdateUser:  projectBizCode.UpdateUser,
	}
}

func ToProjectBizCodeEntity(tProjectBizCode *dao.TProjectBizCode) *entity.ProjectBizCode {
	return &entity.ProjectBizCode{
		ID:          tProjectBizCode.ID,
		BizKey:      tProjectBizCode.BizKey,
		BizValue:    tProjectBizCode.BizValue,
		BizType:     tProjectBizCode.BizType,
		Project:     tProjectBizCode.Project,
		Cluster:     tProjectBizCode.Cluster,
		Service:     tProjectBizCode.Service,
		SpanAlias:   tProjectBizCode.SpanAlias,
		CreatedUser: tProjectBizCode.CreatedUser,
		UpdateUser:  tProjectBizCode.UpdateUser,
	}
}
