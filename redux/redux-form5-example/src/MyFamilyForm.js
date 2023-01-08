import { ButtonToolbar, Col, Form, FormGroup, Button } from "react-bootstrap";
import { Field, FieldArray, reduxForm } from "redux-form";
import asyncValidate from "./asyncValidation";
import myValidation from "./MyValidation";

const renderFamilyFields = (families, index, fields) => (
  <li key={index} className={'list-group-item'}>
    <Button type='button' style={{marginRight: '10px'}} onClick={() => fields.remove(index)}>削除</Button>
    <span style={{marginRight: '10px'}}>家族 {index + 1}</span>
    <Field
      name={`${families}.name`}
      type='text'
      component='input'
      label='家族氏名'
    />
  </li>
);

const renderFamilies = ({ fields }) => (
  <FormGroup>
    <Col componentClass={'familyclass'} sm={2}>家族</Col>
    <Col sm={8}>
      <ButtonToolbar>
        <Button type='button' onClick={() => fields.push({})}>追加</Button>
      </ButtonToolbar>
      <ul className='list-group'>
        {fields.map(renderFamilyFields)}
      </ul>
    </Col>
  </FormGroup>
)

let MyFamilyForm = props => {
  const { handleSubmit, pristine, reset, submitting, contactValue } = props;
  return (
    <Form horizontal onSubmit={handleSubmit}>
      <FieldArray name='families' component={renderFamilies} />
      <FormGroup>
        <Col smoffset={2} sm={5}>
          <ButtonToolbar>
            <Button bsstyle={'primary'} type='submit' disabled={pristine}>登録</Button>
          </ButtonToolbar>
        </Col>
      </FormGroup>
    </Form>
  )
}

MyFamilyForm = reduxForm({
  form: 'myForm',
  destroyOnUnmount: false,
  forceUnregisterOnUnmount: true,
  validate: myValidation,
  asyncValidate,
})(MyFamilyForm);

export default MyFamilyForm;