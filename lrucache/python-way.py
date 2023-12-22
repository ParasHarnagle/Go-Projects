from collections import OrderedDict

class LRUCache:
    def __init__(self,size: int):
    self.cache = OrderedDict()
    self.size = size

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        self.cache.move_to_end(key)
        return self.cache[key]
    
    def put(self, key: int, value:int) -> None:
        #insert item to cache
        if key in self.cache:
            #remove the exisitng key
            del self.cache[key]
        
        #check if cache is full remove fist item from list
        elif len(self.cache) >= self.size:
            self.cache.popitem(last=False)
        #insert key 
        self.cache[key] = value

    def get_MRU(self) -> int:
        if self.cache:
            return next(reversed(self.cache))
        return -1