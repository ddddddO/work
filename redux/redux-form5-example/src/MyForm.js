import { Field, FieldArray, reduxForm, formValueSelector } from "redux-form";
import { connect } from 'react-redux';

import {
  Form,
  FormGroup,
  Col,
  // ControlLabel,
  Button,
  ButtonToolbar,
} from 'react-bootstrap';

import myValidation from "./MyValidation";
import asyncValidate from "./asyncValidation";

const renderField = 
({
  input,
  label,
  type,
  placeholder,
  meta: {touched, error, warning}
}) => {
  const validationState = error ? 'error' : warning ? 'warning' : 'success';
  // HelpBlockはもう無さそうなので、pタグで代替
  return (
    <FormGroup controlId={input.name} validationstate={touched ? validationState : null}>
      <Col componentclass={'temp-label'} sm={2}>{label}</Col>
      <Col sm={5}>
        <input {...input} id={input.name} placeholder={placeholder} type={type} className={'aaa'}/>
        {
          touched && error &&
          <p>{error}</p>
        }
      </Col>
    </FormGroup>
  )
}

let MyForm = props => {
  const { handleSubmit, pristine, reset, submitting, contactValue } = props;
  return (
    // NOTE: before renderField
    //
    // <Form horizontal onSubmit={handleSubmit}>
    //   <FormGroup controlId={'name'}>
    //     <Col componentClass={'temp-label'} sm={2}>お名前</Col>
    //     <Col sm={5}>
    //       <Field
    //         id={'namef'}
    //         name='name'
    //         component='input'
    //         type='text'
    //         placeholder="Name"
    //         className={'form-control'}
    //       />
    //     </Col>
    //   </FormGroup>
    //   <FormGroup>
    //     <Col smOffset={2} sm={5}>
    //       <ButtonToolbar>
    //         <Button bsStyle={'primary'} type='submit' disabled={pristine || submitting}>登録</Button>
    //         <Button type='button' disabled={pristine || submitting} onClick={reset}>クリア</Button>
    //       </ButtonToolbar>
    //     </Col>
    //   </FormGroup>
    // </Form>

    <Form horizontal onSubmit={handleSubmit}>
      <Field
        name='primary_name'
        component={renderField}
        type='text'
        label='お名前'
        placeholder='ウルトラ・マン'
      />
      <FormGroup controlId={'contact'}>
        <Col componentClass={'tttemp'} sm={2}>連絡先</Col>
        <Col sm={5}>
          <label classNam='radio-inline'>
            <Field 
              name='contact'
              id='contact'
              component='input'
              type='radio'
              value='email'
            /> メール
          </label>
          <label classNam='radio-inline'>
            <Field 
              name='contact'
              id='contact'
              component='input'
              type='radio'
              value='phone'
            /> 電話
          </label>
        </Col>
      </FormGroup>
      {
        contactValue === 'email' &&
        <Field 
          name='email'
          component={renderField}
          type='text'
          label='メールアドレス'
          placeholder='aaa@aaa.com'
        />
      }
      {
        contactValue === 'phone' &&
        <Field 
          name='phone'
          component={renderField}
          type='text'
          label='電話番号'
          placeholder='08022223333'
        />
      }
      <FormGroup>
        <Col smoffset={2} sm={5}>
          <ButtonToolbar>
            <Button bsstyle={'primary'} type='submit' disabled={pristine}>次へ</Button>
            <Button type='button' disabled={pristine || submitting} onClick={reset}>クリア</Button>
          </ButtonToolbar>
        </Col>
      </FormGroup>
    </Form>
  );
};

// export default reduxForm({
//   form: 'myForm',
//   validate: myValidation,
//   initialValues: { name: 'だれですか' }, // リロード後、一瞬表示されるが消える
// })(MyForm);

MyForm = reduxForm({
  form: 'myForm',

  destroyOnUnmount: false,
  forceUnregisterOnUnmount: true,

  validate: myValidation,
  initialValues: { primary_name: 'だれですか' },

  asyncValidate,
  asyncBlurFields: ['phone'],
})(MyForm);

const selector = formValueSelector('myForm');

MyForm = connect(state => {
  const contactValue = selector(state, 'contact')
  return {
    contactValue,
  }
})(MyForm)

export default MyForm;