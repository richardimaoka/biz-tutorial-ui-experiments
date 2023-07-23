import { css } from "@emotion/react";
import { useRouter } from "next/router";
import { queryParamToString } from "../../libs/queryString";

export const AutoPlayButton = () => {
  const router = useRouter();
  const autoPlay = queryParamToString(router.query.autoPlay) === "true";

  const startAutoPlay = () => {
    router.replace({ query: { ...router.query, autoPlay: true } });
  };
  const stopAutoPlay = () => {
    router.replace({ query: { ...router.query, autoPlay: undefined } });
  };

  const AutoPlayText = () => (
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

  const AutoPlayStopText = () => (
    <div
      css={css`
        font-size: 16px;
        height: 18px;
      `}
    >
      Stop AutoPlay
    </div>
  );

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
      onClick={autoPlay ? stopAutoPlay : startAutoPlay}
    >
      {autoPlay ? <AutoPlayStopText /> : <AutoPlayText />}
    </button>
  );
};
