/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  purge: ["./views/**/*.html"],
  content: [],
  theme: {
    extend: {}
  },
  plugins: []
}

//npx tailwindcss -i ./src/input.css -o ./public/mystyle.css --watch
