const TikTok = require("./tiktok")

;(async () => {
   const client = new TikTok();
   await client.generateDevice();
   await client.registerDevice();
   return;
   const profile = await client.registerAccount("email@domain.com", "my-password");
   console.log(profile);
})()
