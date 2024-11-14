const form = document.querySelector("main > form");

const password = document.getElementById("password");
const passwordConfirmation = document.getElementById("password-repeat");

form.addEventListener("submit", (event) => {
  if (!form.checkValidity()) {
    event.preventDefault();
    event.stopPropagation();
  }

  form.classList.add("was-validated");
});

password.addEventListener("input", () => {
  console.log("hola");
  if (password.value.length < 8) {
    password.setCustomValidity("Debe tener más de 8 caracteres");
  } else {
    password.setCustomValidity("");
  }
});

passwordConfirmation.addEventListener("input", () => {
  if (password.value !== passwordConfirmation.value) {
    passwordConfirmation.setCustomValidity("Las contraseñas no coinciden");
  } else {
    passwordConfirmation.setCustomValidity("");
  }
});
