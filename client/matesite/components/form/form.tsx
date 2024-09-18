import React, { useState, useEffect, FormEvent } from "react";

interface Issue {
  id: string;
  title: string;
  description: string;
}

interface IssueFormProps {
  onSave: (issue: Issue) => void;
  currentIssue: Issue | null;
  onCancelEdit: () => void;
}

const IssueForm: React.FC<IssueFormProps> = ({
  onSave,
  currentIssue,
  onCancelEdit,
}) => {
  const [title, setTitle] = useState<string>("");
  const [description, setDescription] = useState<string>("");

  useEffect(() => {
    if (currentIssue) {
      setTitle(currentIssue.title);
      setDescription(currentIssue.description);
    }
  }, [currentIssue]);

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    onSave({
      id: currentIssue ? currentIssue.id : Date.now().toString(),
      title,
      description,
    });
    setTitle("");
    setDescription("");
  };

  return (
    <form
      onSubmit={handleSubmit}
      className='space-y-4 bg-white p-6 rounded-lg shadow-md'
    >
      <h2 className='text-xl font-bold text-gray-800'>
        {currentIssue ? "Edit Issue" : "Add Issue"}
      </h2>
      <input
        type='text'
        placeholder='Title'
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        required
        className='w-full p-2 border border-gray-300 rounded text-gray-800'
      />
      <textarea
        placeholder='Description'
        value={description}
        onChange={(e) => setDescription(e.target.value)}
        required
        className='w-full p-2 border border-gray-300 rounded text-gray-800'
      />
      <div className='flex space-x-4'>
        <button
          type='submit'
          className='px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600'
        >
          {currentIssue ? "Update" : "Add"}
        </button>
        {currentIssue && (
          <button
            type='button'
            onClick={onCancelEdit}
            className='px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600'
          >
            Cancel Edit
          </button>
        )}
      </div>
    </form>
  );
};

export default IssueForm;
