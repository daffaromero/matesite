import React from "react";

interface Issue {
  id: string;
  title: string;
  description: string;
}

interface IssueItemProps {
  issue: Issue;
  onEdit: (issue: Issue) => void;
  onDelete: (id: string) => void;
}

const IssueItem: React.FC<IssueItemProps> = ({ issue, onEdit, onDelete }) => {
  return (
    <div className='p-4 border border-gray-300 rounded mb-4 bg-white shadow-sm'>
      <h3 className='text-lg font-bold text-gray-800'>{issue.title}</h3>
      <p className='mb-2 text-gray-600'>{issue.description}</p>
      <button
        onClick={() => onEdit(issue)}
        className='px-4 py-2 bg-yellow-500 text-white rounded mr-2 hover:bg-yellow-600'
      >
        Edit
      </button>
      <button
        onClick={() => onDelete(issue.id)}
        className='px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600'
      >
        Delete
      </button>
    </div>
  );
};

export default IssueItem;
