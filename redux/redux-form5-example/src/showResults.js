// サーバエミュレート用
const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

export default (async function showResults(v) {
  await sleep(500);
  window.alert(`You submitted:\n\n${JSON.stringify(v, null, 2)}`)
});