module.exports = function (ctx) {
  function getFromEnvOrDefault (key, value) {
    if (process.env[key]) {
      return process.env[key]
    }
    return value
  }

  return {
    boot: [
      'bus',
      'components',
      'services'
    ],

    css: [
      'app.css'
    ],

    extras: [
      'fontawesome-v5',
      'roboto-font', // optional, you are not bound to it
      'material-icons' // optional, you are not bound to it
    ],

    framework: {
      iconSet: 'fontawesome-v5', // Quasar icon set
      lang: 'en-us', // Quasar language pack
      all: false,

      config: {
        brand: {
          primary: '#434343',
          secondary: '#C4C4C4',
          accent: '#9C27B0',

          dark: '#1d1d1d',

          positive: '#81Bf97',
          negative: '#DF6756',
          info: '#4FA9D2',
          warning: '#F0DD5D'
        }
      },

      components: [
        'QLayout',
        'QHeader',
        'QDrawer',
        'QPageContainer',
        'QPage',
        'QBtn',
        'QList',
        'QItem',
        'QItemSection',
        'QItemLabel',
        'QDialog',
        'QChip',
        'QCard',
        'QCardSection',
        'QCardActions',
        'QSeparator',
        'QSpace',
        'QToolbar',
        'QToolbarTitle',
        'QMenu',
        'QBadge',
        'QTabs',
        'QTab',
        'QRouteTab',
        'QTabPanels',
        'QTabPanel',
        'QBar'
      ],

      directives: [
        'Ripple'
      ],

      plugins: [
        'Loading',
        'Dialog',
        'Notify'
      ]
    },

    supportIE: false,

    htmlVariables: {
      buildDate: new Date().toISOString(),
      version: process.env.VERSION || '0.0.1'
    },

    build: {
      publicPath: process.env.PUBLIC_PATH || '/',
      env: {
        API_URL: getFromEnvOrDefault('API_URL', '/jsonfiles')
      },
      scopeHoisting: true,
      vueRouterMode: 'history',
      extendWebpack (cfg) {
        cfg.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /node_modules/,
          options: {
            formatter: require('eslint').CLIEngine.getFormatter('stylish')
          }
        })
      }
    },

    devServer: {
      // https: true,
      port: 15235,
      open: false // opens browser window automatically
    },

    // animations: 'all', // --- includes all animations
    animations: [],

    ssr: {
      pwa: false
    },

    pwa: {
      // workboxPluginMode: 'InjectManifest',
      // workboxOptions: {}, // only for NON InjectManifest
      workboxOptions: {
        skipWaiting: true,
        clientsClaim: true
      },
      manifest: {
        name: 'TagHub',
        short_name: 'TagHub',
        description: 'All repositories that I starred in GitHub since 2009 organized by tags',
        display: 'standalone',
        orientation: 'portrait',
        background_color: '#ffffff',
        theme_color: '#434343',
        icons: [
          {
            'src': 'public/icons/icon-128x128.png',
            'sizes': '128x128',
            'type': 'image/png'
          },
          {
            'src': 'public/icons/icon-192x192.png',
            'sizes': '192x192',
            'type': 'image/png'
          },
          {
            'src': 'public/icons/icon-256x256.png',
            'sizes': '256x256',
            'type': 'image/png'
          },
          {
            'src': 'public/icons/icon-384x384.png',
            'sizes': '384x384',
            'type': 'image/png'
          },
          {
            'src': 'public/icons/icon-512x512.png',
            'sizes': '512x512',
            'type': 'image/png'
          }
        ]
      }
    },

    cordova: {
    },

    electron: {
      extendWebpack (cfg) {
      },
      packager: {
      },
      builder: {
      }
    }
  }
}
