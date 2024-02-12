const join = document.querySelector(".join"),
  overlay = document.querySelector(".overlay"),
  closeBtn = document.querySelector(".overlay .close");

  const create = document.querySelector(".createpost"),
  post = document.querySelector(".overlayposts"),
  closepost = document.querySelector(".overlayposts .close");

  const create2 = document.querySelector(".createcategory"),
  catogary = document.querySelector(".overlaycatogaries"),
  closecatogary = document.querySelector(".overlaycatogaries .close");

join.addEventListener("click", () => {
  overlay.classList.add("active");
});

closeBtn.addEventListener("click", () => { 
  overlay.classList.remove("active");
});

window.addEventListener("click", (e) => { 
  if (e.target === overlay){
    overlay.classList.remove("active");
  }
});

create.addEventListener("click", () => {
  post.classList.add("active");
});
closepost.addEventListener("click", () => {
  post.classList.remove("active");
});
window.addEventListener("click", (e) => { 
  if (e.target === catogary){
    catogary.classList.remove("active");
  }
});
create2.addEventListener("click", () => {
  catogary.classList.add("active");
});
closecatogary.addEventListener("click", () => {
  catogary.classList.remove("active");
});
window.addEventListener("click", (e) => { 
  if (e.target === post){
    post.classList.remove("active");
  }
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