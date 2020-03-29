import { Dialog } from 'quasar'

export default class {
  error (err) {
    let message = err
    if (typeof (err) !== 'string') {
      message = err.response ? err.response.data.error : err.toString()
    }
    Dialog.create({
      title: 'Ooops. Something was wrong',
      position: 'bottom',
      message
    })
  }
}
