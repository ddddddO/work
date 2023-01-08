// 非同期でサーバ側バリデーションを模した処理
const sleep = ms => new Promise(resolve => setTimeout(resolve, ms))

export default (async function asyncValidate(values) {
  await sleep(1000);
  if (!!values.phone && !values.phone.match(/^(0[5-9]0[0-9]{8}|0[1-9][1-9][0-9]{7})$/)) {
    throw { phone: '不正な電話番号です' }
  }
});
