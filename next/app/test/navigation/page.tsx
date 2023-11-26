"use client";
import { PrevButton } from "@/app/components/navigation/PrevButton";
import { NextButton } from "@/app/components/navigation/NextButton";
import { PlayButton } from "@/app/components/navigation/PlayButton";
import { PauseButton } from "@/app/components/navigation/PauseButton";

let i = 1;

export default function Page() {
  i++;
  return (
    <div>
      <div>
        {/* <NextButton href={`/test/navigation?param=${i}`} /> */}
        <PrevButton href={`/test/navigation?param=${i}`} />
      </div>
      <div>
        <PlayButton onClick={() => {}} />
        <PauseButton onClick={() => {}} />
      </div>
    </div>
  );
}
