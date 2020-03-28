export default class {
  constructor (deps) {
    this.$http = deps.http
  }

  async readAll () {
    const resp = await this.$http.get('/tags.json')
    return resp.data
  }
}
