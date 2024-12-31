import NextStep from "../components/NextStep";
import { PrintRules } from "../components/PrintRules";

export default function Step1() {
  return (
    <>
      <main className="flex min-h-screen flex-col items-center bg-gradient-to-b from-[#2e026d] to-[#15162c] text-white">
        <div className="mt-10">
          <h1 className="text-4xl">Round 1</h1>
          <div className="mt-10 text-xl">
            <PrintRules label="Deal 10 Cards" />
            <PrintRules label="Two Sets" />
          </div>
          <div className="mt-10">
            <NextStep label="Step 2" link="/step2" classes="" />
          </div>
        </div>
      </main>
    </>
  );
}
