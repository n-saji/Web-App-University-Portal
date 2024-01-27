async function createAccount() {
  let name = document.getElementById("form-name").value;
  let email = document.getElementById("form-email").value;
  let password = document.getElementById("form-password").value;
  let response_form = document.getElementById("response-from-server");
  let next_button = document.getElementById("next-button");
  console.log(next_button.innerHTML);
  if (next_button.innerHTML == "Login") {
    setTimeout(changeToLogin, 3000);
    return;
  }
  if (!name || !email || !password) {
    alert("Empty fields!");
    return;
  }
  let response = await fetch(`http://3.111.149.112:5050/create-account`, {
    method: "POST",
    body: JSON.stringify({
      name: name,
      info: {
        credentials: {
          email_id: email,
          password: password,
        },
      },
    }),
  });

  if (response.status != 500) {
    let responseJSON = await response.json();
    console.log(responseJSON);
    response_form.innerHTML = responseJSON;
    response_form.style.display = "block";
    response_form.style.fontSize = "large";
    next_button.innerHTML = "Login";
  } else {
    let responseJSON = await response.json();
    console.log(responseJSON);
    response_form.innerHTML = responseJSON;
    response_form.style.display = "block";
    response_form.classList.add("error");
    setTimeout(removeElement, 2000);
  }
}

function removeElement() {
  let response_form = document.getElementById("response-from-server");
  response_form.style.display = "none";
}

function changeToLogin() {
  window.location.replace("index.html");
}

function setindex() {
  window.location.replace("index.html");
}
