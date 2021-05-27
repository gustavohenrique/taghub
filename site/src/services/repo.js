export default class {
  constructor (deps) {
    this.$http = deps.http
  }

  async searchByTag (id) {
    const resp = await this.$http.get(`/tags/${id}.json`)
    const data = resp.data || {}
    return {
      items: data.data || [],
      total: data.meta ? data.meta.total : 0
    }
  }

  async search (filter) {
    const { page } = filter.pagination
    const resp = await this.$http.get(`/repos/${page}.json`)
    const data = resp.data || {}
    return {
      items: data.data || [],
      total: data.meta ? data.meta.total : 0
    }
  }
}
