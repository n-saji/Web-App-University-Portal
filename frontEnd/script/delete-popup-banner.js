let indexForDeletePopUp1, indexForDeletePopUp2, indexForDeletePopUp3;
function deletePopupTrue(index1, index2, index3) {
  let popupbanner = document.querySelector(".delete_popup");
  popupbanner.classList.remove("delete_popup_close");
  popupbanner.classList.add("delete_popup_display");
  indexForDeletePopUp1 = index1;
  indexForDeletePopUp2 = index2;
  indexForDeletePopUp3 = index3;
  window.onkeydown = function (event) {
    if (event.keyCode == 27) {
      deleteFalse();
    }
  };
}
function deleteFalse() {
  let popupbanner = document.querySelector(".delete_popup");
  popupbanner.classList.add("delete_popup_close");
  let msg = document.querySelector(".error-message");
  msg.innerHTML = "";
}


