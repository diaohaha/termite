import request from '../utils/request';

export const fetchFlows = (query) => {
    return request({
        url: '/api/config/flow/query/',
        method: 'post',
        data: query
    })
}

export const queryFlowConfigs = (query) => {
	return request({
		url: '/api/config/flow/query/v2/',
		method: 'post',
		data: query
	})
}

export const updateFlowSwitch = (query) => {
	return request({
		url: '/api/config/flow/switch/',
		method: 'post',
		data: query
	})
}

export const fetchFlowInstances = (query) => {
	return request({
		url: '/api/instance/flow/query/',
		method: 'post',
		data: query
	})
}

export const addFlowInstances = (query) => {
	return request({
		url: '/api/instance/flow/add/',
		method: 'post',
		data: query
	})
}

export const addFlowInstancesByFile = (query) => {
	// return service.service.post('/api/instance/flow/add/file/', query)
	return request({
		headers: {
			'Content-Type': 'multipart/form-data'
		},
		url: '/api/instance/flow/add/file/',
		method: 'post',
		data: query
	})
}

export const recoverFlowInstances = (query) => {
	return request({
		url: '/api/instance/flow/recover/',
		method: 'post',
		data: query
	})
}

export const deleteFlowInstances = (query) => {
	return request({
		url: '/api/instance/flow/delete/',
		method: 'post',
		data: query
	})
}

export const recoverWorkInstances = (query) => {
	return request({
		url: '/api/instance/work/recover/',
		method: 'post',
		data: query
	})
}

export const fetchWorkInstances = (query) => {
	return request({
		url: '/api/instance/work/query/',
		method: 'post',
		data: query
	})
}
export const fetchWorksConfig = (query) => {
    return request({
        url: '/api/config/work/query/',
        method: 'post',
        data: query
    })
}
export const queryWorkConfigs = (query) => {
	return request({
		url: '/api/config/work/query/v2/',
		method: 'post',
		data: query
	})
}

export const createFlowConfig = (params) => {
	return request({
		url: '/api/config/flow/create/',
		method: 'post',
		data: params
	})
}

export const updateFlowConfig = (params) => {
	return request({
		url: '/api/config/flow/update/',
		method: 'post',
		data: params
	})
}

export const updateWorkConfig = (params) => {
	return request({
		url: '/api/config/work/update/',
		method: 'post',
		data: params
	})
}

export const copyFlowConfig = (params) => {
	return request({
		url: '/api/config/flow/copy/',
		method: 'post',
		data: params
	})
}

export const copyWorkConfig = (params) => {
	return request({
		url: '/api/config/work/copy/',
		method: 'post',
		data: params
	})
}
export const createWorkConfig = (params) => {
	return request({
		url: '/api/config/work/create/',
		method: 'post',
		data: params
	})
}

export const deleteFlowConfig = (params) => {
	return request({
		url: '/api/config/flow/delete/',
		method: 'post',
		data: params
	})
}

export const deleteWorkConfig = (params) => {
	return request({
		url: '/api/config/work/delete/',
		method: 'post',
		data: params
	})
}


// info

export const getFlowCountInfo = (params) => {
	return request({
		url: '/api/info/flow/count/',
		method: 'post',
		data: params
	})
}

export const getWorkCountInfo = (params) => {
	return request({
		url: '/api/info/work/count/',
		method: 'post',
		data: params
	})
}
