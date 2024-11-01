import { Client, Talent } from "@/models/user";

export const mockTalents: Talent[] = [
    {
      id: "",
      name: "Dummy Talent 1",
      email: "dummyuserOne@talent.com",
      phone: "+880123456789",
      is_verified: false,
    },
    {
      id: "",
      name: "Dummy Talent 2",
      email: "dummyuserTwo@talent.com",
      phone: "+880888888888",
      is_verified: true,
    },
];

export const mockClients: Client[] = [
  {
    id: "",
    name: "Dummy Client 1",
    email: "dummyuserOne@client.com",
    phone: "+880123456789",
    is_verified: false,
  },
  {
    id: "",
    name: "Dummy Client 2",
    email: "dummyuserTwo@client.com",
    phone: "+880888888888",
    is_verified: true,
  },
];
