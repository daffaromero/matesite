export const createIssue = async (issue: {
  title: string;
  description: string;
}) => {
  const response = await fetch("http://localhost:8000/issues/new", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ issue }),
  });
  return response.json();
};

export const listIssues = async () => {
  const response = await fetch("http://localhost:8000/issues");
  return response.json();
};

export const getIssue = async (id: string) => {
  const response = await fetch(`http://localhost:8000/issues/${id}`);
  return response.json();
};

export const updateIssue = async (
  id: string,
  issue: { title: string; description: string }
) => {
  const response = await fetch(`http://localhost:8000/issues/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ issue }),
  });
  return response.json();
};

export const deleteIssue = async (id: string) => {
  const response = await fetch(`http://localhost:8000/issues/${id}`, {
    method: "DELETE",
  });
  return response.json();
};
