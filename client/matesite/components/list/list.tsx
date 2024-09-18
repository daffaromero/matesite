import React from "react";
import IssueItem from "../item/item";
import { useIssues } from "../../hooks/useIssues";
import { deleteIssue } from "../../api/issues";

interface Issue {
  id: string;
  title: string;
  description: string;
}

interface IssueListProps {
  onEdit: (issue: Issue) => void;
}

const IssueList: React.FC<IssueListProps> = ({ onEdit }) => {
  const { issues, isLoading, isError, mutate } = useIssues();

  if (isLoading) return <div className='text-gray-800'>Loading...</div>;
  if (isError) return <div className='text-red-500'>Error loading issues</div>;

  const handleDelete = async (id: string) => {
    await deleteIssue(id);
    mutate();
  };

  return (
    <div className='mt-8'>
      <h2 className='text-xl font-bold mb-4 text-gray-800'>Issue List</h2>
      {issues.map((issue: Issue) => (
        <IssueItem
          key={issue.id}
          issue={issue}
          onEdit={onEdit}
          onDelete={handleDelete}
        />
      ))}
    </div>
  );
};

export default IssueList;
