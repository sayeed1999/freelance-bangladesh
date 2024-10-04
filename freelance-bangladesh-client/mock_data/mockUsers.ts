import { Client, Talent } from "@/models/user";

export const mockTalents: Talent[] = [
    {
      ID: "",
      Name: "Dummy Talent 1",
      Email: "dummyuserOne@talent.com",
      Phone: "+880123456789",
      IsVerified: false,
    },
    {
      ID: "",
      Name: "Dummy Talent 2",
      Email: "dummyuserTwo@talent.com",
      Phone: "+880888888888",
      IsVerified: true,
    },
];

export const mockClients: Client[] = [
  {
    ID: "",
    Name: "Dummy Client 1",
    Email: "dummyuserOne@client.com",
    Phone: "+880123456789",
    IsVerified: false,
  },
  {
    ID: "",
    Name: "Dummy Client 2",
    Email: "dummyuserTwo@client.com",
    Phone: "+880888888888",
    IsVerified: true,
  },
];
