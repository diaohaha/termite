
export const lowerJSONKey = (jsonObj) => {
	for (var key in jsonObj){
		console.log(key)
		var keyArr = ["Works", "Dags", "Dependences", "TriggerRule"]
		if (keyArr.indexOf(key) != -1) {
			console.log(key)
			if (key == "TriggerRule") {
				jsonObj["trigger_rule"] = jsonObj[key];
			} else {
				jsonObj[key.toLowerCase()] = jsonObj[key];
			}
			delete(jsonObj[key]);
		}
		if (typeof jsonObj[key.toLowerCase()] == 'object') {
			jsonObj[key.toLowerCase()] = lowerJSONKey(jsonObj[key.toLowerCase()]);
		}
	}
	return jsonObj;
};
