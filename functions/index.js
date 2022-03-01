const functions = require("firebase-functions");

// // Create and Deploy Your First Cloud Functions
// // https://firebase.google.com/docs/functions/write-firebase-functions
//
// exports.helloWorld = functions.https.onRequest((request, response) => {
//   functions.logger.info("Hello logs!", {structuredData: true});
//   response.send("Hello from Firebase!");
// });

admin.initializeApp();
const db = admin.firestore();
const GHcollection = db.collection("github");

// Fetch repository information from GitHub
exports.GitHubToDB = functions.pubsub.schedule('0 * * * *')
  .timeZone('America/New_York').onRun(async () => {
    const response = await fetch('https://api.github.com/graphql', {
      method: 'POST',
      headers: {
        Authorization: `bearer ${functions.config().github.key}`
      },
      body: JSON.stringify({
        query: `
          {
            user(login: "luuk180") {
              repositories(orderBy: {field: PUSHED_AT, direction: ASC}, first: 100) {
                edges {
                  node {
                    isPrivate
                    name
                    url
                    homepageUrl
                    description
                    diskUsage
                  }
                }
              }
            }
          }
          `,
      }),
    });
    const json = await response.json();
    const jsonObject = JSON.stringify(json.data.user.repositories.edges);
    for(let i = 0; i < JSON.parse(jsonObject).length; i++) {
      const currObj = JSON.parse(jsonObject)[i];
      if(!currObj.node.isPrivate) {
        await GHcollection.doc(currObj.node.name).set(currObj.node);
      }
    }

  return 0;
});