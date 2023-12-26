const join = document.querySelector(".join"),
  overlay = document.querySelector(".overlay"),
  closeBtn = document.querySelector(".overlay .close");

join.addEventListener("click", () => {
  overlay.classList.add("active");
});

closeBtn.addEventListener("click", () => {
  overlay.classList.remove("active");
});

function togglePass() {
  var x = document.getElementById("passIn");
  var txt = document.getElementById("toggleTxt");
  if (x.type === "password") {
    x.type = "text";
    txt.textContent = "Hide";
  } else {
    x.type = "password";
    txt.textContent = "Show";
  }
}

// const catOptions = document.getElementById('catOptions');
// if (catOptions) {
//   const lastOption = catOptions.querySelector('option:last-child');
//   if (lastOption) {
//     lastOption.innerHTML = "{{ .CategoryName}}";
//   }
// }