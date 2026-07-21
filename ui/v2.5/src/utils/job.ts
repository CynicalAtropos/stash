import { useCallback, useEffect, useRef, useState } from "react";
import { getWSClient, useWSState } from "src/core/StashService";
import {
  Job,
  JobStatus,
  JobStatusUpdateType,
  useFindJobQuery,
  useJobsSubscribeSubscription,
} from "src/core/generated-graphql";

export type JobFragment = Pick<
  Job,
  "id" | "status" | "subTasks" | "description" | "progress" | "error"
>;

export const useMonitorJob = (
  jobID: string | undefined | null,
  onComplete?: (job?: JobFragment) => void
) => {
  const { state } = useWSState(getWSClient());

  const jobsSubscribe = useJobsSubscribeSubscription({
    skip: !jobID,
  });
  const {
    data: jobData,
    loading,
    startPolling,
    stopPolling,
  } = useFindJobQuery({
    variables: {
      input: { id: jobID ?? "" },
    },
    fetchPolicy: "network-only",
    skip: !jobID,
  });

  const [job, setJob] = useState<JobFragment | undefined>();
  const completedJobID = useRef<string>();
  const onCompleteRef = useRef(onComplete);

  useEffect(() => {
    onCompleteRef.current = onComplete;
  }, [onComplete]);

  useEffect(() => {
    if (!jobID) {
      completedJobID.current = undefined;
    }
  }, [jobID]);

  const completeJob = useCallback(
    (completedJob?: JobFragment) => {
      if (!jobID || completedJobID.current === jobID) {
        return;
      }

      completedJobID.current = jobID;
      stopPolling();
      setJob(undefined);
      onCompleteRef.current?.(completedJob);
    },
    [jobID, stopPolling]
  );

  useEffect(() => {
    if (!jobID) {
      return;
    }

    if (loading) {
      return;
    }

    const j = jobData?.findJob;
    if (j) {
      if (
        j.status === JobStatus.Finished ||
        j.status === JobStatus.Failed ||
        j.status === JobStatus.Cancelled
      ) {
        completeJob(j);
      } else {
        setJob(j);
      }
    } else {
      // must've already finished
      completeJob();
    }
  }, [jobID, jobData, loading, completeJob]);

  // monitor job
  useEffect(() => {
    if (!jobID) {
      return;
    }

    if (!jobsSubscribe.data) {
      return;
    }

    const event = jobsSubscribe.data.jobsSubscribe;
    if (event.job.id !== jobID) {
      return;
    }

    if (event.type !== JobStatusUpdateType.Remove) {
      setJob(event.job);
    } else {
      completeJob(event.job);
    }
  }, [jobsSubscribe, jobID, completeJob]);

  // it's possible that the websocket connection isn't present
  // in that case, we'll just poll the server
  useEffect(() => {
    if (!jobID || completedJobID.current === jobID) {
      stopPolling();
      return;
    }

    if (state === "connected") {
      stopPolling();
    } else {
      const defaultPollInterval = 1000;
      startPolling(defaultPollInterval);
    }
  }, [jobID, state, startPolling, stopPolling]);

  return { job };
};
