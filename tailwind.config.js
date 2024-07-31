/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./web/**/*.{html,templ}'],
    darkMode: false, // or 'media' or 'class'
    theme: {
        extend: {},
    },
    plugins: [require('daisyui'),],
}

