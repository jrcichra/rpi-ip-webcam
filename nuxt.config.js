export default {
    modules: ['bootstrap-vue/nuxt', '@nuxtjs/axios', 'moment', 'nuxt-socket-io',],
    components: true,
    moment: {
        timezone: true
    }, buildModules: [
        '@nuxtjs/fontawesome',
    ],
    server: {
        port: 3000,
        host: '0.0.0.0',
    },
    fontawesome: {
        icons: {
            solid: true,
            regular: true,
        }
    }, plugins: [
        '@/plugins/axios',
    ],
    ssr: false,// Disable Server Side rendering
    io: {
        // module options
        sockets: [{
            name: 'main',
            url: 'http://secpi.pk5001z'
        }]
    }
    // router: {
    //     base: '/dist/'
    // }
}
