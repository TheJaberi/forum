const join = document.querySelector(".join"),
  overlay = document.querySelector(".overlay"),
  closeBtn = document.querySelector(".overlay .close"),
  passToggle = document.querySelector(".passToggle");

join.addEventListener("click", () => {
  overlay.classList.add("active");
});

closeBtn.addEventListener("click", () => {
  overlay.classList.remove("active");
});

passToggle.addEventListener("click", () => {
  var passLbl = document.getElementById("passLbl");
  if (passLbl.textContent == "show")
  {
    passLbl.textContent = "hide";
  }
  else if (passLbl.textContent == "hide")
  {
    passLbl.textContent = "show";
  }
});