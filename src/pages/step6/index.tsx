import NextStep from "../components/NextStep";
import { PrintRules } from "../components/PrintRules";

export default function Step6() {
  return (
    <>
      <main className="flex min-h-screen flex-col items-center bg-gradient-to-b from-[#2e026d] to-[#15162c] text-white">
        <div className="mt-10">
          <h1 className="text-4xl">Round 6</h1>
          <div className="mt-10 text-xl">
            <PrintRules label="Deal 12 Cards" />
            <PrintRules label="1 Set and 2 Sequences" />
          </div>
          <div className="mt-10">
            <NextStep label="Step 7" link="/step7" classes="" />
          </div>
        </div>
      </main>
    </>
  );
}
