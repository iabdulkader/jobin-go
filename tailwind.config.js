/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./**/*.html",
    "./**/*.templ",
    "./**/*.go",
    "./views/**/*.templ",
    "./components/**/*.templ",
  ],
  theme: {
    extend: {
      colors: {
        "custom-gray": "#B99470",
        "custom-body": "#002b36",
      },
      fontFamily: {
        dank: ["dank-mono", "sans-serif"],
      },
    },
  },
  plugins: [],
};
