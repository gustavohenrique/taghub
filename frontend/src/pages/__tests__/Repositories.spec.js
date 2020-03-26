/* eslint-disable */
import { mount, shallowMount } from '@vue/test-utils'
import flushPromises from 'flush-promises'
import Repositories from '../Repositories.vue'


const localVue = global.localVue
const mountWithMocks = mocks => {
  return shallowMount(Repositories, {
    localVue,
    mocks
  })
}

describe('pages/Repositories', () => {
  let pagination = null

  beforeEach(() => {
    pagination = {
      page: 1,
      prevPage: 0,
      rowsPerPage: 10,
      info: {}
    }
  })

  it('fetch items when mounted', async () => {
    const result = {
      pagination: {
        page: 1,
        prevPage: 1,
        rowsPerPage: 30,
        info: {
          startCursor: 'Y3Vyc29yOnYyOpHOAB9hZQ==',
          endCursor: 'Y3Vyc29yOnYyOpHOAB9haw==',
          hasNextPage: true,
          hasPreviousPage: false
        }
      },
      items: [
        {
          node: {
            name: 'my-repo',
            description: 'a mocked repo',
            owner: {
              login: 'gustavohenrique'
            }
          }
        }
      ]
    }

    const mocked = jest.fn({ pagination }).mockResolvedValue(result)
    const mocks = {
      $s: {
        repository: { fetchItems: mocked }
      }
    }
    const wrapper = mountWithMocks(mocks)
    await flushPromises()
    expect(mocked).toBeCalled();

    const { vm } = wrapper
    expect(vm.pagination).toEqual(result.pagination)
    expect(vm.items).toEqual(result.items)
  })

  it('should router to form page', async () => {
    const $router = []
    const mocks = {
      $router
    }
    const wrapper = mountWithMocks(mocks)
    const owner = 'gustavohenrique'
    const name = 'backend'
    wrapper.vm.edit(owner, name)
    expect($router).toHaveLength(1)
    expect($router[0]).toEqual({
      name: 'form',
      params: {
        name,
        owner
      }
    })
  })
})

