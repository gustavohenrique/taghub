import axios from 'axios'

const instance = axios.create({
  baseURL: process.env.API_URL,
  timeout: 120000
})

instance.defaults.headers.common['Content-Type'] = 'application/json'

const deps = {
  http: instance,
  axios: axios
}

const requireFile = require.context(
  '../services',
  false,
  /[\w-]+\.js$/
)

const services = {}
requireFile.keys().forEach(fileName => {
  const config = requireFile(fileName)
  const name = fileName
    .replace(/^\.\//, '')
    .replace(/^\.\/_/, '')
    .replace(/\.\w+$/, '')
  const Service = config.default || config
  services[name] = new Service(deps)
})

export default ({ Vue }) => {
  Vue.prototype.$s = services
}
