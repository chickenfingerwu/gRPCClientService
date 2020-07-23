import 'protobufjs'
import axios from 'axios'

const baseURL = 'http://127.0.0.1:8081'

const getAll = () => {
    console.log("called getAll")
    const request = axios.get(baseURL + '/customers/2')
    return request.then(response =>
        {
            console.log(response.data)
            return response.data
        })
}

export default {
    getAll: getAll
}