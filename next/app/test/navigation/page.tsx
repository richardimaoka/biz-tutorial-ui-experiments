import { PrevButton } from "@/app/components/navigation2/PrevButton";
import { NextButton } from "@/app/components/navigation2/NextButton";

let i = 1;

export default async function Page() {
  i++;
  return (
    <div>
      <NextButton href={`/test/navigation?param=${i}`} />
      <PrevButton href={`/test/navigation?param=${i}`} />
    </div>
  );
}
