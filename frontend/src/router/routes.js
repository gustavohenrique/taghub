
const routes = [
  {
    path: '/',
    component: () => import('layouts/Layout.vue'),
    children: [
      { path: 'sync', name: 'sync', component: () => import('pages/Repos/Sync.vue') },
      { path: 'repositories', name: 'repositories', component: () => import('pages/Repos/List.vue') }
    ]
  }
]

// Always leave this as last one
if (process.env.MODE !== 'ssr') {
  routes.push({
    path: '*',
    component: () => import('pages/Error404.vue')
  })
}

export default routes
