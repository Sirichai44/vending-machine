/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  darkMode: 'class',
  theme: {
    extend: {
      fontFamily: {
        nurito: ['Nurito', 'sans-serif'],
        popins: ['Poppins', 'sans-serif'],
        comfortaa: ['Comfortaa', 'cursive']
      },
      colors: {
        'primary-light': '#374151',
        'primary-dark': '#D1D5DB'
      }
    }
  },
  plugins: [],
  corePlugins: {
    transitionProperty: true
  }
};