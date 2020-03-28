<template>
  <q-layout view="hhh LpR fFf" class="bg-grey-1">

    <q-header
      elevated
      class="text-white"
      style="background: #344150"
      height-hint="61.59"
    >
      <q-toolbar class="q-py-sm q-px-md">
        <div class="q-pt-xs">
          <logo color="white" width="140" />
        </div>

        <q-space />

        <div
          class="gt-md q-ml-xs q-gutter-md text-body2 text-weight-bold row items-center no-wrap"
        >
          <q-btn
            v-for="m in menus"
            :key="JSON.stringify(m)"
            :label="m.label"
            :to="m.route"
            flat
            text-color="white"
            no-caps
            type="a"
            class="text-md"
          />
        </div>

        <div class="gt-md q-pl-sm q-gutter-sm row items-center no-wrap">
          <q-btn round dense flat :ripple="false" :icon="fabGithub" size="19px" color="white" no-caps />
        </div>

        <div class="lt-md">
          <q-btn
            flat
            dense
            round
            @click="isVisible = !isVisible"
            aria-label="Menu"
            icon="menu"
          >
            <q-menu
              anchor="bottom right"
              self="top right"
            >
              <q-list
                bordered
                separator
                style="min-width: 100px"
              >
                <q-item
                  clickable
                  :to="item.route"
                  v-for="item in menus"
                  :key="JSON.stringify(item)"
                >
                  <q-item-section>
                    <q-item-label>{{ item.label }}</q-item-label>
                    <q-item-label caption>{{ item.description }}</q-item-label>
                  </q-item-section>
                  <q-separator />
                </q-item>
              </q-list>
            </q-menu>
          </q-btn>
        </div>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { fabGithub } from '@quasar/extras/fontawesome-v5'

export default {
  data () {
    return {
      fabGithub,
      isVisible: true,
      text: '',
      options: null,
      filteredOptions: []
    }
  },
  computed: {
    menus () {
      return [
        {
          label: 'Repositories',
          description: 'List all starred repositories',
          route: { name: 'repositories' }
        },
        {
          label: 'Tags',
          description: 'Delete or rename your tags',
          route: { name: 'tags' }
        },
        {
          label: 'Synchronization',
          description: 'Fetch new starred repositories from GitHub',
          route: { name: 'sync' }
        }
      ]
    }
  }
}
</script>
