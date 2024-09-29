import React from "react";

interface FormType {
  formTitle: string;
  formItems: FormItemType[];
  dispatchAction: () => any;
  submitBtnName: string;
}

interface FormItemType {
  label: string;
  name: string;
  type: string;
  id: string;
  ref: React.MutableRefObject<any>;
  placeholder: string;
  required: boolean;
  validationError: string;
}

const Form = (form: FormType) => {
  const handleSubmitEvent = (e: any): any => {
    e.preventDefault();
    // Call API using dispatch action
    form.dispatchAction();
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
              ref={formItem.ref}
              placeholder={formItem.placeholder}
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
              ref={formItem.ref}
              // checked={!!formItem.value}
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
              ref={formItem.ref}
              placeholder={formItem.placeholder}
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
