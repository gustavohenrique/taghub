import Vue from 'vue'
const bus = new Vue()

const b = {
  publish: (evt, ...args) => bus.$emit(evt, ...args),
  unsubscribe: (evt, ...args) => bus.$off(evt, ...args),
  subscribe: (evt, callback) => {
    if (typeof evt === 'string') {
      bus.$on(evt, callback)
    }
    if (Array.isArray(evt)) {
      for (const e of evt) {
        bus.$on(e, callback)
      }
    }
  }
}

Vue.prototype.$publish = b.publish
Vue.prototype.$subscribe = b.subscribe
Vue.prototype.$unsubscribe = b.unsubscribe
