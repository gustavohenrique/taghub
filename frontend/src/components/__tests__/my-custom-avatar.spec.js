/* eslint-disable */
import { mount, shallowMount } from '@vue/test-utils'
import MyCustomAvatar from '../my-custom-avatar.vue'

/**
 * Tipos de assert:
 * https://jestjs.io/docs/en/expect
 * .toEqual(value)
 * .toBe(value)
 * .toHaveBeenCalled()
 * .toHaveBeenCalledTimes(number)
 * .toHaveBeenCalledWith(arg1, arg2, ...)
 * .toHaveReturned()
 * .toHaveReturnedTimes(number)
 * .toHaveLength(number)
 * .toHaveProperty(keyPath, value?)
 * .toBeCloseTo(number, numDigits?)
 * .toBeDefined()
 * .toBeTruthy()
 * .toBeFalsy()
 * .toBeGreaterThan(number)
 * .toBeLessThan(number)
 * .toBeNull()
 * .toBeUndefined()
 * .toContain(item)
 * .toMatch(regexpOrString)
 * .toThrow(error?)
 */

describe('MyCustomAvatar', () => {
  const localVue = global.localVue
  const item = {
    avatarUrl: 'https://my.avatar.com/me.jpg',
    login: 'gustavohenrique'
  }
  const wrapper = mount(MyCustomAvatar, {
    localVue,
    propsData: { item }
  })
  const vm = wrapper.vm

  it('is a valid Vue component', () => {
    expect(wrapper.isVueInstance()).toBe(true)
  })

  it('contains the login', () => {
    expect(wrapper.find('span').text()).toContain(item.login)
  })
})
