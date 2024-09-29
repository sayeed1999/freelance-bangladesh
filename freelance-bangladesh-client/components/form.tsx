import React from "react";
import { useEffect, useState } from "react";

interface FormType {
  formTitle: string;
  formItems: FormItemType[];
  dispatchAction: (payload: any) => any;
  submitBtnName: string;
}

interface FormItemType {
  label: string;
  name: string;
  value: string | number | readonly string[] | undefined;
  type: string;
  id: string;
  placeholder: string;
  required: boolean;
  validationError: string;
}

const Form = (form: FormType) => {
  const [input, setInput] = useState<any>({});

  const handleSubmitEvent = (e: any): any => {
    e.preventDefault();

    const requiredFields = form.formItems.filter((x) => x.required);

    Object.keys(input).forEach((key) => {
      // Note: If it is a required field but not provided in payload, then throw validation error!
      const requiredField = requiredFields.find(
        (requiredField) => requiredField.name === key
      );

      if (requiredField && !input[key]) {
        return alert(`Please Provide a Valid ${requiredField.label}`);
      }
    });

    // Call API using dispatch action
    form.dispatchAction(input);
  };

  const handleInput = (e: any) => {
    if (!e.target) return;

    const { name, value } = e.target;

    setInput((prev: any) => ({
      ...prev,
      [name]: value,
    }));
  };

  const renderFormItem = (formItem: FormItemType, index: number) => {
    switch (formItem.type) {
      case "textarea":
        return (
          <div key={index} className="mb-4">
            <label className="block text-gray-700">{formItem.label}</label>
            <textarea
              id={formItem.id}
              name={formItem.name}
              value={input[formItem.name] ?? formItem.value}
              placeholder={formItem.placeholder}
              onChange={handleInput}
              className="w-full px-4 py-2 border rounded-md"
              rows={4}
            />
            <div id={formItem.id} className="sr-only">
              {formItem.validationError}
            </div>
          </div>
        );
      case "checkbox":
        return (
          <div key={index} className="mb-4">
            <label className="block text-gray-700">{formItem.label}</label>
            <input
              type="checkbox"
              id={formItem.id}
              name={formItem.name}
              checked={input[formItem.name] ?? formItem.value}
              onChange={handleInput}
              className="mr-2"
            />
            {formItem.label}
            <div id={formItem.id} className="sr-only">
              {formItem.validationError}
            </div>
          </div>
        );
      default:
        return (
          <div key={index} className="mb-4">
            <label className="block text-gray-700">{formItem.label}</label>
            <input
              type={formItem.type ?? "text"}
              id={formItem.id}
              name={formItem.name}
              value={input[formItem.name] ?? formItem.value}
              placeholder={formItem.placeholder}
              onChange={handleInput}
              className="w-full px-4 py-2 border rounded-md"
            />
            <div id={formItem.id} className="sr-only">
              {formItem.validationError}
            </div>
          </div>
        );
    }
  };

  return (
    <div className="container mx-auto py-16">
      <h2 className="text-3xl font-bold text-center">{form.formTitle}</h2>
      <form onSubmit={handleSubmitEvent} className="mt-8 max-w-md mx-auto">
        {form.formItems.map((formItem, index) =>
          renderFormItem(formItem, index)
        )}
        <button className="btn-submit bg-blue-600 text-white px-4 py-2 rounded-md w-full">
          {form.submitBtnName}
        </button>
      </form>
    </div>
  );
};

export default Form;
