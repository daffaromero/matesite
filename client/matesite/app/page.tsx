"use client";

import React, { useState } from "react";
import IssueForm from "@/components/form/form";
import IssueList from "@/components/list/list";
import { createIssue, updateIssue } from "@/api/issues";
import { useIssues } from "@/hooks/useIssues";

interface Issue {
  id: string;
  title: string;
  description: string;
}

const App: React.FC = () => {
  const { mutate } = useIssues();
  const [currentIssue, setCurrentIssue] = useState<Issue | null>(null);

  const handleSave = async (issue: Issue) => {
    if (currentIssue) {
      await updateIssue(currentIssue.id, {
        title: issue.title,
        description: issue.description,
      });
    } else {
      await createIssue({ title: issue.title, description: issue.description });
    }
    mutate();
    setCurrentIssue(null);
  };

  const handleEdit = (issue: Issue) => {
    setCurrentIssue(issue);
  };

  const handleCancelEdit = () => {
    setCurrentIssue(null);
  };

  return (
    <div className='p-8 bg-gray-100 min-h-screen'>
      <h1 className='text-2xl font-bold mb-8 text-gray-800'>Issues CRUD</h1>
      <div className='max-w-2xl mx-auto'>
        <IssueForm
          onSave={handleSave}
          currentIssue={currentIssue}
          onCancelEdit={handleCancelEdit}
        />
        <IssueList onEdit={handleEdit} />
      </div>
    </div>
  );
};

export default App;
