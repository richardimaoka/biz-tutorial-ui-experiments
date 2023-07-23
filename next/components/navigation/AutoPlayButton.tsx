import { css } from "@emotion/react";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

interface Scheduling {
  kind: "Scheduling";
}

interface Scheduled {
  kind: "Scheduled";
  timeoutId: number;
}

interface Stopping {
  kind: "Stopping";
  timeoutId: number;
}

interface Stopped {
  kind: "Stopped";
}

type AutoPlayState = Scheduling | Scheduled | Stopping | Stopped;

interface AutoPlayButtonProps {
  nextStep: string;
}

export const AutoPlayButton = ({ nextStep }: AutoPlayButtonProps) => {
  const [state, setState] = useState<AutoPlayState>({ kind: "Stopped" });
  const router = useRouter();

  // effectful code
  useEffect(() => {
    switch (state.kind) {
      case "Scheduling":
        const tid = window.setTimeout(() => {
          router.replace({ query: { ...router.query, step: nextStep } });
          setState({ kind: "Scheduling" }); // after transitioned to next step, schedule again transition to next next step
        }, 1000);
        setState({ kind: "Scheduled", timeoutId: tid });
        break;
      case "Scheduled":
        break; // do nothing
      case "Stopping":
        window.clearTimeout(state.timeoutId);
        setState({ kind: "Stopped" });
        break;
      case "Stopped":
        break; // do nothing
      default:
        const _exhaustiveCheck: never = state;
        return _exhaustiveCheck;
    }
  }, [state, nextStep, router]);

  const onClick = () => {
    switch (state.kind) {
      case "Scheduling":
        break; // do nothing
      case "Scheduled":
        setState({ kind: "Stopping", timeoutId: state.timeoutId });
        break;
      case "Stopping":
        break; // do nothing
      case "Stopped":
        setState({ kind: "Scheduling" });
        break;
      default:
        const _exhaustiveCheck: never = state;
        return _exhaustiveCheck;
    }
  };

  const AutoPlayText = (): JSX.Element => {
    switch (state.kind) {
      case "Scheduling":
        return <></>;
      case "Scheduled":
        return (
          <div
            css={css`
              font-size: 16px;
              height: 18px;
            `}
          >
            Stop AutoPlay
          </div>
        );
      case "Stopping":
        return <></>;
      case "Stopped":
        return (
          <>
            <div
              css={css`
                font-size: 16px;
                height: 18px;
              `}
            >
              AutoPlay
            </div>
            <div
              css={css`
                font-size: 8px;
                line-height: 8px;
              `}
            >
              to next milestone
            </div>
          </>
        );
    }
  };

  return (
    <button
      css={css`
        //positioning
        position: fixed;
        bottom: 0px;
        left: 50%;
        transform: translate(-50%, 0%);

        //sizing
        width: 120px;
        height: 40px;

        // color and styles
        background-color: rgba(255, 255, 255, 0.8);
        color: black;
        border-style: none;
      `}
      onClick={onClick}
    >
      <AutoPlayText />
    </button>
  );
};
