let globalEmaild,
  globalePassword = "";
let api_response;
let instrId = getCookie("account_id");
let cookie_token = getCookie("token");
async function autofil() {
  let api_url = "";
  let url = `http://3.111.149.112:5050/view-profile-instructor/${instrId}`;

  let api_response = await fetch(url, {
    method: "GET",
    headers: { Token: cookie_token },
  }).catch((err) => {
    api_url = err;
  });
  if (api_url != "") {
    alert("Internal Server Error Please Login Again" + api_url);
    return "";
  }

  let name_input = document.querySelector(".js-input-name");
  let email_input = document.querySelector(".js-input-email");
  let password_input = document.querySelector(".js-input-password");
  let code_input = document.querySelector(".js-input-department");

  let response = await api_response.json();

  name_input.value = response.Name;
  email_input.value = response.Credentials.email_id;
  password_input.value = response.Credentials.password.slice(0, 9);
  code_input.value = response.Code;
  globalEmaild = email_input.value;
  globalePassword = password_input.value;
}

autofil();

async function ToUpdateDetails() {
  let name_input = document.querySelector(".js-input-name");
  let email_input = document.querySelector(".js-input-email");
  let password_input = document.querySelector(".js-input-password");
  let code_input = document.querySelector(".js-input-department");

  console.log(
    email_input.value,
    password_input.value,
    globalEmaild,
    globalePassword
  );

  let api_url = "";
  let url = `http://3.111.149.112:5050/update-instructor?instructor_id=${instrId}`;
  let api_response = await fetch(url, {
    method: "PATCH",
    headers: { Token: cookie_token },
    body: JSON.stringify({
      instructor_name: name_input.value,
      instructor_code: code_input.value,
    }),
  }).catch((err) => {
    api_url = err;
  });
  if (api_url != "") {
    alert("Internal Server Error updating basic- " + api_url);
    return "";
  }

  if (
    globalEmaild !== email_input.value &&
    globalePassword === password_input.value
  ) {
    let api_url = "";
    let url = `http://3.111.149.112:5050/update-instructor-credentials`;
    api_response = await fetch(url, {
      method: "PUT",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        id: instrId,
        email_id: email_input.value,
      }),
    }).catch((err) => {
      api_url = err;
    });
    if (api_url != "") {
      alert("Internal Server Error updating email-" + api_url);
      return "";
    }
  }
  if (
    globalePassword !== password_input.value &&
    globalEmaild === email_input.value
  ) {
    let api_url = "";
    let url = `http://3.111.149.112:5050/update-instructor-credentials`;
    api_response = await fetch(url, {
      method: "PUT",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        id: instrId,
        password: password_input.value,
      }),
    }).catch((err) => {
      api_url = err;
    });
    if (api_url != "") {
      alert("Internal Server Error updating password-" + api_url);
      return "";
    }
  }
  if (
    globalePassword !== password_input.value &&
    globalEmaild !== email_input.value
  ) {
    let api_url = "";
    let url = `http://3.111.149.112:5050/update-instructor-credentials`;
    api_response = await fetch(url, {
      method: "PUT",
      headers: { Token: cookie_token },
      body: JSON.stringify({
        id: instrId,
        password: password_input.value,
        email_id: email_input.value,
      }),
    }).catch((err) => {
      api_url = err;
    });
    if (api_url != "") {
      alert("Internal Server Error updating both-" + api_url);
      return "";
    }
  }
  globalEmaild = email_input.value;
  globalePassword = password_input.value;
  if (api_response.status != 500) {
    window.location.reload();
  }
}
