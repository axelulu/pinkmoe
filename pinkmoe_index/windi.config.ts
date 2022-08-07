import { defineConfig } from 'vite-plugin-windicss';

export default defineConfig({
  darkMode: 'class',
  theme: {
    extend: {
      keyframes: {
        jello: {
          '0%': {
            transform: 'none',
          },
          '11.1%': {
            transform: 'none',
          },
          to: {
            transform: 'none',
          },
          '22.2%': {
            transform: 'skewX(-12.5deg) skewY(-12.5deg)',
          },

          '33.3%': {
            transform: 'skewX(6.25deg) skewY(6.25deg)',
          },

          '44.4%': {
            transform: 'skewX(-3.125deg) skewY(-3.125deg)',
          },

          '55.5%': {
            transform: 'skewX(1.5625deg) skewY(1.5625deg)',
          },

          '66.6%': {
            transform: 'skewX(-.78125deg) skewY(-.78125deg)',
          },

          '77.7%': {
            transform: 'skewX(.390625deg) skewY(.390625deg)',
          },

          '88.8%': {
            transform: 'skewX(-.1953125deg) skewY(-.1953125deg)',
          },
        },
        rotating: {
          from: { transform: 'rotate(0)' },
          to: { transform: 'rotate(360deg)' },
        },
        pinkZoomInUp: {
          '0%': {
            opacity: 0,
            transform: 'scale3d(.9, .9, 1) translateY(10%)',
          },
          to: {
            opacity: 1,
            transform: 'scaleX(1)',
          },
        },
        pinkZoomOutDown: {
          '0%': {
            opacity: 1,
            transform: 'scaleX(1)',
          },
          to: {
            opacity: 0,
            transform: 'scale3d(.9, .9, 1) translateY(10%)',
          },
        },
        pinkTipZoomInUp: {
          '0%': {
            opacity: 0,
            transform: 'scale3d(.95, .95, 1) translateY(2%)',
          },
          to: {
            opacity: 1,
            transform: 'scaleX(1)',
          },
        },
        pinkTipZoomOutDown: {
          '0%': {
            opacity: 1,
            transform: 'scaleX(1)',
          },
          to: {
            opacity: 0,
            transform: 'scale3d(.95, .95, 1) translateY(2%)',
          },
        },
        fadeIn: {
          '0%': { opacity: '0', display: 'none' },
          to: { opacity: '1', display: 'block' },
        },
        fadeOut30: {
          '0%': { opacity: '1', display: 'block' },
          to: { opacity: '0', display: 'none' },
        },
        swing: {
          '20%': { transform: 'rotate3d(0, 0, 1, 15deg)', transformOrigin: 'top center' },
          '40%': { transform: 'rotate3d(0, 0, 1, 10deg)', transformOrigin: 'top center' },
          '60%': { transform: 'rotate3d(0, 0, 1, 5deg)', transformOrigin: 'top center' },
          '80%': { transform: 'rotate3d(0, 0, 1, -5deg)', transformOrigin: 'top center' },
          to: { transform: 'rotate3d(0, 0, 1, 0deg)', transformOrigin: 'top center' },
        },
        lazy_blur: {
          '0%': {
            filter: 'blur(5px)',
          },
          '100%': {
            filter: 'blur(0)',
          },
        },
        lazy_back_blur: {
          '0%': {
            'backdrop-filter': 'blur(8px)',
          },
          '100%': {
            'backdrop-filter': 'blur(0)',
          },
        },
      },
      animation: {
        fadeIn30: 'fadeIn 0.3s normal 1 forwards',
        fadeOut30: 'fadeOut30 0.3s normal 1 forwards',
        pinkZoomInUp30: 'pinkZoomInUp 0.3s normal 1 forwards',
        pinkZoomOutDown30: 'pinkZoomOutDown 0.3s normal 1 forwards',
        pinkTipZoomInUp30: 'pinkTipZoomInUp 0.3s normal 1 forwards',
        pinkTipZoomOutDown30: 'pinkTipZoomOutDown 0.3s normal 1 forwards',
        lazyloaded: 'lazy_blur 0.8s',
        spinlazyloaded: 'lazy_back_blur 0.5s',
        swing30: 'swing 0.3s normal 2',
        rotating: 'rotating 1.2s linear infinite',
        jello100: 'jello 1s normal 1 forwards',
      },
      boxShadow: {
        scollTop: '0 0 10px 0 rgb(0 0 0 / 15%)',
        menuFixed: '0 0 10px 0 rgb(0 0 0 / 15%)',
        menu: '0 0 10px 0 rgb(0 0 0 / 15%)',
        shadowGlobal: '0 0 0 2px rgb(253 242 248/var(--tw-bg-opacity))',
      },
      opacity: {
        '0!': '0!important',
      },
    },
  },
});
