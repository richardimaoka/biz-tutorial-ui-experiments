"use client";
import { PrevButton } from "@/app/components/navigation2/PrevButton";
import { NextButton } from "@/app/components/navigation2/NextButton";
import { PlayButton } from "@/app/components/navigation2/PlayButton";
import { PauseButton } from "@/app/components/navigation2/PauseButton";

let i = 1;

export default function Page() {
  i++;
  return (
    <div>
      <div>
        <div>
          <PrevButton href={`/test/navigation?param=${i}`} />
        </div>
      </div>
      <div>
        <PlayButton onClick={() => {}} />
        <PauseButton onClick={() => {}} />
      </div>
    </div>
  );
}
