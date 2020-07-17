import 'protobufjs'
import axios from 'axios'

const baseURL = 'http://127.0.0.1:9090'

const getAll = () => {
    console.log("called getAll")
    const request = axios.get(baseURL + '/customers')
    return request.then(response =>
        {
            console.log(response.data)
            return response.data
        })
}

export default {
    getAll: getAll
}