import useSWR from "swr";
import { listIssues } from '@/api/issues';

const fetcher = (url: string) => listIssues();

export const useIssues = () => {
  const { data, error, mutate } = useSWR(
    "http://localhost:8000/issues",
    fetcher
  );

  return {
    issues: data?.issues,
    isLoading: !error && !data,
    isError: error,
    mutate,
  };
};
