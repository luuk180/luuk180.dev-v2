<template>
  <div id="home-page" class="text-gray-700 text-xl">
    <h1 class="text-center">{{welcome_text}}</h1>
    This is going to be my portfolio and I'm very busy building it right now!
  </div>
</template>

<script>
let actualtext = "Welcome to luuk180.dev!";
let textfield = "";
for(let i = 0; i > actualtext.length; i++) {
  setTimeout(() => {textfield.appendChild(actualtext[i]);}, 1000);
  welcome_text = textfield;
}
</script>