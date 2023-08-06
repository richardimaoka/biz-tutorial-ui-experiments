"use client";

import { useRouter, useSearchParams } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

import styles from "./style.module.css";

interface Scheduled {
  kind: "Scheduled";
  timeoutId: number;
}

interface Transitioned {
  kind: "Transitioned";
  step: string;
}

interface Stopped {
  kind: "Stopped";
}

type AutoPlayState = Scheduled | Transitioned | Stopped;

interface AutoPlayButtonProps {
  nextStep: string;
}

export const AutoPlayButton = ({ nextStep }: AutoPlayButtonProps) => {
  const [state, setState] = useState<AutoPlayState>({ kind: "Stopped" });
  const router = useRouter();
  const searchParams = useSearchParams();

  // effectful code
  useEffect(() => {
    let newParams = new URLSearchParams();
    searchParams.forEach((value, key) => {
      newParams.set(key, value);
    });
    newParams.set("step", nextStep);

    switch (state.kind) {
      case "Scheduled":
        break; // do nothing
      case "Transitioned":
        // This is an important state as React re-renders this component *BEFORE* updating the URL query string
        // without the Transitioned state, setTimeout is called twice for the same nextStep
        if (state.step !== nextStep) {
          const tid = window.setTimeout(() => {
            router.push("/?" + newParams.toString());
            setState({ kind: "Transitioned", step: nextStep });
          }, 1000);
          setState({ kind: "Scheduled", timeoutId: tid });
        }
        break;
      case "Stopped":
        break; // do nothing
      default:
        const _exhaustiveCheck: never = state;
        return _exhaustiveCheck;
    }
  }, [state, nextStep, router, searchParams]);

  const onClick = () => {
    let newParams = new URLSearchParams();
    searchParams.forEach((value, key) => {
      newParams.set(key, value);
    });
    newParams.set("step", nextStep);

    switch (state.kind) {
      case "Scheduled":
        window.clearTimeout(state.timeoutId);
        setState({ kind: "Stopped" });
        break;
      case "Transitioned":
        setState({ kind: "Stopped" });
        break;
      case "Stopped":
        router.push("/?" + newParams.toString());
        setState({ kind: "Transitioned", step: nextStep });
        break;
      default:
        const _exhaustiveCheck: never = state;
        return _exhaustiveCheck;
    }
  };

  const AutoPlayText = (): JSX.Element => {
    switch (state.kind) {
      case "Scheduled": // fallthrough
      case "Transitioned":
        return <div className={styles.text}>Stop AutoPlay</div>;
      case "Stopped":
        return (
          <>
            <div className={styles.text}>AutoPlay</div>
            <div className={styles.minitext}>to next milestone</div>
          </>
        );
    }
  };

  return (
    <button className={styles.autoplay} onClick={onClick}>
      <AutoPlayText />
    </button>
  );
};
