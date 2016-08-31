import platform from './platform'

const engines = {
    // cms,
    platform
};

export default {
    routes () {
        return Object.keys(engines).reduce(function (obj, en) {
            return obj.concat(engines[en].routes)
        }, [])
    },
    reducers () {
        return Object.keys(engines).reduce(function (obj, en) {
            return Object.assign(obj, engines[en].reducers)
        }, {})
    }

}