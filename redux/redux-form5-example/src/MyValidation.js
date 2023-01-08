const myValidation = vs => {
  const errors = {}
  if (!vs.name) {
    errors.name = '必須項目です'
  } else if (vs.name.length > 5) {
    errors.name = '5文字以内で指定してください'
  }
  return errors;
}

export default myValidation;