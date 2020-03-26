export const events = {
  LOADING_START: 'loading_start',
  LOADING_STOP: 'loading_stop',
  REQUEST_ERROR: 'request_error',
  DIALOG_ERROR: 'dialog_error_visible',

  CREATED: 'created',
  SAVED: 'saved',
  CHANGED: 'changed',
  EDIT: 'edit',
  BACK_TO_LIST: 'back_to_list',
  TALK_TO_SEARCHBOX: 'talk_to_searchbox',

  SET_COMPANY: 'set_comapny',
  AUTOCOMPLETE_PRODUCT: 'autocomplete_product',
  AUTOCOMPLETE_STUDENT: 'autocomplete_student',
  AUTOCOMPLETE_CONTENT: 'autocomplete_content',
  SET_SELLING_STARTS_AT: 'selling_set_starts_at',
  SET_SELLING_ENDS_AT: 'selling_set_ends_at'
}

export const TOKEN = 'token'

export default {
  events,
  TOKEN
}
