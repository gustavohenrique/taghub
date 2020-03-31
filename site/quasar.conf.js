module.exports = function (ctx) {
  function getFromEnvOrDefault (key, value) {
    if (process.env[key]) {
      return JSON.stringify(process.env[key])
    }
    return JSON.stringify(value)
  }
  const API_URL = getFromEnvOrDefault('API_URL', 'statics/jsonfiles')

  return {
    boot: [
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
        'QBadge'
      ],

      directives: [
        'Ripple'
      ],

      plugins: [
        'Loading',
        'Dialog'
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
        API_URL
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
      manifest: {
        // name: 'Quasar App',
        // short_name: 'Quasar App',
        // description: 'A Quasar Framework app',
        display: 'standalone',
        orientation: 'portrait',
        background_color: '#ffffff',
        theme_color: '#027be3',
        icons: [
          {
            'src': 'statics/icons/icon-128x128.png',
            'sizes': '128x128',
            'type': 'image/png'
          },
          {
            'src': 'statics/icons/icon-192x192.png',
            'sizes': '192x192',
            'type': 'image/png'
          },
          {
            'src': 'statics/icons/icon-256x256.png',
            'sizes': '256x256',
            'type': 'image/png'
          },
          {
            'src': 'statics/icons/icon-384x384.png',
            'sizes': '384x384',
            'type': 'image/png'
          },
          {
            'src': 'statics/icons/icon-512x512.png',
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
